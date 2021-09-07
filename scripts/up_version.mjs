#!/usr/bin/env zx

$.verbose = false

try {
    // Fetch the version declaration from the version.go file
    const versionDeclaration = await $`cat config/version.go | grep VERSION`

    // Since the Go syntax is of the form const VERSION = "x.x.x" and this is also valid js
    // We can eval to get the version string
    const version = eval(`
      (() => {
          ${versionDeclaration}
          
          return VERSION
        }
      )()
    `)

    console.info(`--- Removing version.go file for version ${version}`)

    await $`rm config/version.go`

    const newVersion = await question(`--- What should the new version be? `)

    console.info(`--- Creating version.go file for version ${newVersion}`)

    const newVersionFile = `package config

const VERSION = "${newVersion}"`

    await $`echo ${newVersionFile} >> config/version.go`

    console.info("--- Pushing new version to Github")

    await $`git tag v${newVersion} && git push origin v${newVersion}`

    console.info(`--- Version ${newVersion} successfully published`)
} catch (p) {
  console.error(`Exit code: ${p.exitCode}`)
  console.error(`Error: ${p.stderr}`)
}

