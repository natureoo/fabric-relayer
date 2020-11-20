/*
* Copyright (C) 2020 The poly network Authors
* This file is part of The poly network library.
*
* The poly network is free software: you can redistribute it and/or modify
* it under the terms of the GNU Lesser General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* The poly network is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU Lesser General Public License for more details.
* You should have received a copy of the GNU Lesser General Public License
* along with The poly network . If not, see <http://www.gnu.org/licenses/>.
 */
package main

import (
	"strings"

	"github.com/polynetwork/fabric-relayer/config"
	"github.com/urfave/cli"
)

var (
	logLevelFlag = cli.UintFlag{
		Name:  "loglevel",
		Usage: "Set the log level to `<level>` (0~6). 0:Trace 1:Debug 2:Info 3:Warn 4:Error 5:Fatal 6:MaxLevel",
		Value: config.DEFAULT_LOG_LEVEL,
	}

	configPathFlag = cli.StringFlag{
		Name:  "cliconfig",
		Usage: "Server config file `<path>`",
		Value: config.DEFAULT_CONFIG_FILE_NAME,
	}

	logDirFlag = cli.StringFlag{
		Name:  "logdir",
		Usage: "log directory",
		Value: "./Log/",
	}

	polyStartHeight = cli.Uint64Flag{
		Name: "poly",
		Usage: "the poly start height",
		Value: 0,
	}

	fabStartHeight = cli.Uint64Flag{
		Name: "fabric",
		Usage: "the fabric start height",
		Value: 3,
	}
)

//getFlagName deal with short flag, and return the flag name whether flag name have short name
func getFlagName(flag cli.Flag) string {
	name := flag.GetName()
	if name == "" {
		return ""
	}
	return strings.TrimSpace(strings.Split(name, ",")[0])
}
