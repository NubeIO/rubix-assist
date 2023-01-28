# CHANGELOG

## [v0.8.5](https://github.com/NubeIO/rubix-assist/tree/v0.8.5) (2023-01-28)

- Remove the existing mappings for old deployment support

## [v0.8.5](https://github.com/NubeIO/rubix-assist/tree/v0.8.5) (2023-01-25)

- Add OpenVPN feature

## [v0.8.4](https://github.com/NubeIO/rubix-assist/tree/v0.8.4) (2023-01-11)

- Change underscore headers to dash headers as a normal convention

## [v0.8.3](https://github.com/NubeIO/rubix-assist/tree/v0.8.3) (2022-12-21)

- Fix: read file API params

## [v0.8.2](https://github.com/NubeIO/rubix-assist/tree/v0.8.2) (2022-12-12)

- Remove suffix slash (/) from APIs for to support reverse proxy

## [v0.8.1](https://github.com/NubeIO/rubix-assist/tree/v0.8.1) (2022-12-12)

- Add root-dir on systemctl service file creation
- Set ubuntu-18.04 as the runner OS & update packages

## [v0.8.0](https://github.com/NubeIO/rubix-assist/tree/v0.8.0) (2022-12-04)

- Improvement on GetAppStatus
    - Differentiate whether it'd valid request or not
- Upgrade list plugins API
- Fix: delete older installation & listing plugin failure
- Add APIs for plugins

## [v0.7.0](https://github.com/NubeIO/rubix-assist/tree/v0.7.0) (2022-11-24)

- Misc changes to decouple rubix-edge from rubix-assist
- Improvement on plugin/app installation

## [v0.6.2](https://github.com/NubeIO/rubix-assist/tree/v0.6.2) (2022-11-14)

- Remove assist, ffclient, wires CLI from here & add it on rubix-ui

## [v0.6.1](https://github.com/NubeIO/rubix-assist/tree/v0.6.1) (2022-11-13)

- CLI creation on creating of host
    - otherwise, old host details can take place
- Fix: NewForce CLI port issue for bios

## [v0.6.0](https://github.com/NubeIO/rubix-assist/tree/v0.6.0) (2022-11-13)

- Add BIOS implementation for rubix-edge installation
- Add token implementations

## [v0.5.5](https://github.com/NubeIO/rubix-assist/tree/v0.5.5) (2022-10-26)

- Added edge proxy

## [v0.5.4](https://github.com/NubeIO/rubix-assist/tree/v0.5.4) (2022-10-24)

- Remove External token check on proxy.go

## [v0.5.3](https://github.com/NubeIO/rubix-assist/tree/v0.5.3) (2022-10-21)

- Fix: test cases

## [v0.5.2](https://github.com/NubeIO/rubix-assist/tree/v0.5.2) (2022-10-20)

- Integration of edge-bios
- Ongoing APIs creation for edge-bios
- CLI improvements
- Rename typo assitcli to assistcli

## [v0.5.1](https://github.com/NubeIO/rubix-assist/tree/v0.5.1) (2022-10-13)

- FormatRestyV2Response for detecting connection issue
- Change app store location to support different arch modes

## [v0.5.0](https://github.com/NubeIO/rubix-assist/tree/v0.5.0) (2022-09-22)

- Lots of improvements

## [v0.4.6](https://github.com/NubeIO/rubix-assist/tree/v0.4.6) (2022-08-22)

- updates to get networking

## [v0.4.4](https://github.com/NubeIO/rubix-assist/tree/v0.4.4) (2022-08-22)

- added new edge api

## [v0.4.3](https://github.com/NubeIO/rubix-assist/tree/v0.4.3) (2022-08-22)

- fix bug on host ip

## [v0.4.1](https://github.com/NubeIO/rubix-assist/tree/v0.4.1) (2022-08-22)

- add ff proxy
- added basic client for the users
- edge client for date, time and networking

## [v0.3.8](https://github.com/NubeIO/rubix-assist/tree/v0.3.8) (2022-08-19)

- add auth to apis

## [v0.2.7](https://github.com/NubeIO/rubix-assist/tree/v0.2.7) (2022-08-12)

- add plugins apis

## [v0.2.6](https://github.com/NubeIO/rubix-assist/tree/v0.2.5) (2022-08-12)

- Bump installer and dirs version
- added api to remove an app

## [v0.2.5](https://github.com/NubeIO/rubix-assist/tree/v0.2.5) (2022-08-11)

- Bump installer version
- Change product schema

## [v0.2.4](https://github.com/NubeIO/rubix-assist/tree/v0.2.4) (2022-08-11)

- Added edge app installer

## [v0.2.3](https://github.com/NubeIO/rubix-assist/tree/v0.2.3) (2022-08-10)

- List apps
- Fix: port issue

## [v0.2.2](https://github.com/NubeIO/rubix-assist/tree/v0.2.2) (2022-08-09)

- First initial release for rubix-service installable
