package main

import (
	"context"
	"fmt"
	"github.com/blacktop/lporg/database"
	"github.com/blacktop/lporg/dock"
	"github.com/jinzhu/gorm"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/apex/log"
	"github.com/blacktop/lporg/database/utils"
	"github.com/pkg/errors"
)

var porg = `
                                          '.:/+ooossoo+/:-'
                                     ':+ydNMMMMMMMMMMMMMMMNmyo:'
                                   '.--.''.:ohNMMMMMMMMNho:.''..'
                                        -o+.  ':sNMMms-'  .--
                                 '+o     .mNo    '::'   :dNh'    '+-
                                :mMo      dMM-         .NMMs      hNs
                               -NMMNs:--/hMMM+         .MMMNo-''-sMMMs
                          -'   sMMMMMMNNMMMMM:          hMMMMNNNNMMMMN
                         -y    /MMMMMMMMMMMNs           .mMMMMMMMMMMMd     :
                        .mN.    oNMMMMMMMms-             .yNMMMMMMMNy.    -N:
                        hMMm+'   ./syys+-'    -//   ':/-   ./syyys/.    'oNMm'
                       /MMMMMms:.              '.    '''             ./ymMMMMs
                       mMMMMMMMMNo                 '               'sNMMMMMMMN.
                      -MMMMMMMMN+'             :shmmmh+.            'oNMMMMMMMo
                      /MMMMMMMd-             :dds+///sdm/             :mMMMMMMm
                      sMMMMMMd.             :m+'      ':hs'            -mMMMMMM+
                    .hMMMMMMM-             :h:'         'oh.            /MMMMMMN/
                   /mMMMMMMMN'           ./+' .://:::::.  /d/           .MMMMMMMN/
                 'sMMMMMMMMMM/       '.-:-'     '....'     .so-'        oMMMMMMMMN:
                .hMMMMMMMMMMMNs-'     ''                     '--     '-yNMMMMMMMMMN-
               -dMMMMMMMMMMMMMMNs'                                  'yNMMMMMMMMMMMMm.
              -mMMMMMMMMMMMMMMNo'                                    'sMMMMMMMMMMMMMd'
             :NMMMMMMMMMMMMMMm:                                        /NMMMMMMMMMMMMy
            -NMMMMMMMMMMMMMMd.                                        ' -mMMMMMMMMMMMMo
           .mMMMMMMMMMMMMMMd-/o'                                      .o::mMMMMMMMMMMMN/
          'dMMMMMMMMMMMMMMMhmm.                                        -mdhMMMMMMMMMMMMm.
          yMMMMMMMMMMMMMMMMMN-                                          :NMMMMMMMMMMMMMMh
         /MMMMMMMMMMMMMMMMMN:                                            +MMMMMMMMMMMMMMM+
        'mMMMMMMMMMMMMMMMMMo                                              sMMMMMMMMMMMMMMN.
        oMMMMMMMMMMMMMMMMMh                                               'dMMMMMMMMMMMMMMy
       'mMMMMMMMMMMMMMMMMN.                                                :MMMMMMMMMMMMMMM-
       :MMMMMMMMMMMMMMMMMo                                                  yMMMMMMMMMMMMMMy
       sMMMMMMMMMMMMMMMMN'                                                  -MMMMMMMMMMMMMMN'
       dMMMMMMMMMMMMMMMMh                                                    mMMMMMMMMMMMMMM-
       mMMMMMMMMMMMMMMMMo                                                    yMMMMMMMMMMMMMM-
       mMMMMMMMMMMMMMMMMo                                                    sMMMMMMMMMMMMMM.
       hMMMMMMMMMMMMMMMMs                                                    hMMMMMMMMMMMMMN
       oMMMMMMMMMMMMMMMMy                                                    dMMMMMMMMMMMMMy
       .MMMMMMMMMMMMMMMMd                                                    NMMMMMMMMMMMMM-
        yMMMMMMMMMMMMMMMM'                                                  .MMMMMMMMMMMMMh
        .NMMMMMMMMMMMMMMM/                                                  oMMMMMMMMMMMMN-
         :NMMMMMMMMMMMMMMh                                                  mMMMMMMMMMMMMo
          /NMMMMMMMMMMMMMM-                                                /MMMMMMMMMMMMd'
           :NMMMMMMMMMMMMMh                                               'mMMMMMMMMMMMm.
            .hMMMMMMMMMMMMM/                                              oMMMMMMMMMMMN-
              +mMMMMMMMMMMMN-                                            :MMMMMMMMMMMm-
               'oNMMMMMMMMMMm.                                          -NMMMMMMMMMMd-
                 .omNmh+:hNMMm-                                        :NNsmMMMMMMMy'
                   '.     -smMN+                                     'oNh- 'sNMMNh:
                            ':yNh-                                  -hh:     .:-'
                               ':o/'                              '/+.
                                   '                              '

`

var appHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}
Author:{{if .Author}} {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

// RunCommand runs cmd on file
func RunCommand(ctx context.Context, cmd string, args ...string) (string, error) {

	var c *exec.Cmd

	if ctx != nil {
		c = exec.CommandContext(ctx, cmd, args...)
	} else {
		c = exec.Command(cmd, args...)
	}

	output, err := c.Output()
	if err != nil {
		return string(output), err
	}

	// check for exec context timeout
	if ctx != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("command %s timed out", cmd)
		}
	}

	return string(output), nil
}

func restartDock() error {
	ctx := context.Background()

	utils.Indent(log.Info)("restarting Dock")
	if _, err := RunCommand(ctx, "killall", "Dock"); err != nil {
		return errors.Wrap(err, "killing Dock process failed")
	}

	// let system settle
	time.Sleep(5 * time.Second)

	return nil
}

func removeOldDatabaseFiles(dbpath string) error {

	paths := []string{
		filepath.Join(dbpath, "db"),
		filepath.Join(dbpath, "db-shm"),
		filepath.Join(dbpath, "db-wal"),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			utils.DoubleIndent(log.WithField("path", path).Warn)("file not found")
			continue
		}
		if err := os.Remove(path); err != nil {
			return errors.Wrap(err, "removing file failed")
		}
		utils.DoubleIndent(log.WithField("path", path).Info)("removed old file")

	}

	return restartDock()
}

func savePath(confPath string, icloud bool) string {

	if icloud {

		iCloudPath, err := getiCloudDrivePath()
		if err != nil {
			log.WithError(err).Fatal("get iCloud drive path failed")
		}

		if len(confPath) > 0 {
			return filepath.Join(iCloudPath, confPath)
		}

		host, err := os.Hostname()
		if err != nil {
			log.WithError(err).Fatal("get hostname failed")
		}
		host = strings.TrimRight(host, ".local")

		return filepath.Join(iCloudPath, ".launchpad."+host+".yaml")
	}

	if len(confPath) > 0 {
		return confPath
	}

	// get current user
	user, err := user.Current()
	if err != nil {
		log.WithError(err).Fatal("get current user failed")
	}

	return filepath.Join(user.HomeDir, ".launchpad.yaml")
}

func getiCloudDrivePath() (string, error) {

	// get current user
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return filepath.Join(user.HomeDir, "Library/Mobile Documents/com~apple~CloudDocs"), nil
}

func getLaunchpadConfig(verbose bool) (database.Config, error) {
	var (
		launchpadRoot int
		dashboardRoot int
		items         []database.Item
		dbinfo        []database.DBInfo
		conf          database.Config
	)

	// find launchpad database
	tmpDir := os.Getenv("TMPDIR")
	lpad.Folder = filepath.Join(tmpDir, "../0/com.apple.dock.launchpad/db")
	lpad.File = filepath.Join(lpad.Folder, "db")
	// lpad.File = "./launchpad.db"
	if _, err := os.Stat(lpad.File); os.IsNotExist(err) {
		utils.Indent(log.WithError(err).WithField("path", lpad.File).Fatal)("launchpad DB not found")
	}
	utils.Indent(log.WithFields(log.Fields{"database": lpad.File}).Info)("found launchpad database")

	// open launchpad database
	db, err := gorm.Open("sqlite3", lpad.File)
	if err != nil {
		return conf, err
	}
	defer db.Close()

	if verbose {
		db.LogMode(true)
	}

	// get launchpad and dashboard roots
	if err := db.Where("key in (?)", []string{"launchpad_root", "dashboard_root"}).Find(&dbinfo).Error; err != nil {
		log.WithError(err).Error("dbinfo query failed")
	}
	for _, info := range dbinfo {
		switch info.Key {
		case "launchpad_root":
			launchpadRoot, _ = strconv.Atoi(info.Value)
		case "dashboard_root":
			dashboardRoot, _ = strconv.Atoi(info.Value)
		default:
			log.WithField("key", info.Key).Error("bad key")
		}
	}

	// get all the relavent items
	if err := db.Not("uuid in (?)", []string{"ROOTPAGE", "HOLDINGPAGE", "ROOTPAGE_DB", "HOLDINGPAGE_DB", "ROOTPAGE_VERS", "HOLDINGPAGE_VERS"}).
		Order("items.parent_id, items.ordering").
		Find(&items).Error; err != nil {
		log.WithError(err).Error("items query failed")
	}

	// create parent mapping object
	log.Info("collecting launchpad/dashboard pages")
	parentMapping := make(map[int][]database.Item)
	for _, item := range items {
		db.Model(&item).Related(&item.App)
		// db.Model(&item).Related(&item.Widget)
		db.Model(&item).Related(&item.Group)

		parentMapping[item.ParentID] = append(parentMapping[item.ParentID], item)
	}

	log.Info("interating over launchpad pages")
	conf.Apps, err = parsePages(launchpadRoot, parentMapping)
	if err != nil {
		return conf, errors.Wrap(err, "unable to parse launchpad pages")
	}

	log.Info("interating over dashboard pages")
	conf.Widgets, err = parsePages(dashboardRoot, parentMapping)
	if err != nil {
		return conf, errors.Wrap(err, "unable to parse dashboard pages")
	}

	log.Info("interating over dock apps")
	dPlist, err := dock.LoadDockPlist()
	for _, item := range dPlist.PersistentApps {
		conf.DockItems = append(conf.DockItems, item.TileData.FileLabel)
	}
	conf.DockItems = append(conf.DockItems, "============")
	for _, item := range dPlist.PersistentOthers {
		conf.DockItems = append(conf.DockItems, item.TileData.FileLabel)
	}

	return conf, nil
}

func split(buf []string, lim int) [][]string {
	var chunk []string
	chunks := make([][]string, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
