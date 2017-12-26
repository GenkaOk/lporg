# launchpad-organizer :construction: [WIP]

[![Circle CI](https://circleci.com/gh/blacktop/launchpad-organizer.png?style=shield)](https://circleci.com/gh/blacktop/launchpad-organizer) [![GitHub release](https://img.shields.io/github/release/blacktop/launchpad-organizer.svg)](https://github.com/https://github.com/blacktop/launchpad-organizer/releases/releases) [![License](http://img.shields.io/:license-mit-blue.svg)](http://doge.mit-license.org)

> Organize Your macOS Launchpad Apps

---

> **NOTE:** Tested on High Sierra

## Getting Started

```sh
Usage: lporg [OPTIONS] COMMAND [arg...]
Organize Your Launchpad
Version: SNAPSHOT-06629a760be0c6881c7015e6c5a5241df4e3c812, BuildTime: 2017-12-26T19:40:06Z
Author: blacktop - <https://github.com/blacktop>

Options:
  --verbose, -V  verbose output
  --help, -h     show help
  --version, -v  print the version

Commands:
  default  Organize by Categories
  help     Shows a list of commands or help for one command

Run 'lporg COMMAND --help' for more information on a command.
```

```
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
```

```json
{
  "RowID": 122,
  "App": {
    "ItemID": 122,
    "Title": "Spotify",
    "BundleID": "com.spotify.client",
    "StoreID": {
      "String": "",
      "Valid": false
    },
    "CategoryID": {
      "Int64": 8,
      "Valid": true
    },
    "Category": {
      "ID": 8,
      "UTI": "public.app-category.music"
    },
    "Moddate": 527776226,
    "Bookmark": "stuff"
  },
  "UUID": "A5A1AAAA-AAAA-AAAA-AAAA-A055AAAD93B",
  "Flags": {
    "Int64": 0,
    "Valid": true
  },
  "Type": {
    "Int64": 4,
    "Valid": true
  },
  "Group": {
    "ID": 173,
    "CategoryID": {
      "Int64": 0,
      "Valid": false
    },
    "Title": {
      "String": "",
      "Valid": false
    }
  },
  "ParentID": 172,
  "Ordering": {
    "Int64": 29,
    "Valid": true
  }
}
```

## TODO

* [ ] swith to Apex log
* [ ] figure out how to write to DB and not just read :disappointed:
