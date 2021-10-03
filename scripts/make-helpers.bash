#!/usr/bin/env bash

set_current_version() {
  CURRENT_VERSION="$(grep VERSION < config/version.go)"
  IFS='=' read -ra TEMP_VERSION <<< "$CURRENT_VERSION"
  IFS='.' read -ra CURRENT_VERSION <<< "${TEMP_VERSION[1]}"
  CURRENT_VERSION[0]="${CURRENT_VERSION[0]//\"/}"
  CURRENT_VERSION[0]="${CURRENT_VERSION[0]// /}"
  CURRENT_VERSION[2]="${CURRENT_VERSION[2]//\"/}"
}

set_new_version() {
  set_current_version

  NEW_VERSION[0]=${CURRENT_VERSION[0]}
  NEW_VERSION[1]=${CURRENT_VERSION[1]}
  NEW_VERSION[2]=${CURRENT_VERSION[2]}

  case $1 in
    patch)
      # shellcheck disable=SC2003
      # shellcheck disable=SC2086
      NEW_VERSION[2]=$(expr ${CURRENT_VERSION[2]} + 1)
      ;;
    minor)
      # shellcheck disable=SC2003
      # shellcheck disable=SC2086
      NEW_VERSION[1]=$(expr ${CURRENT_VERSION[1]} + 1)
      ;;
    major)
      # shellcheck disable=SC2003
      # shellcheck disable=SC2086
      NEW_VERSION[0]=$(expr ${CURRENT_VERSION[0]} + 1)
      ;;
    *)
      print "unknown version bump, not in (patch, minor, major)"
      exit 1
      ;;
  esac

  export CURRENT_VERSION="${CURRENT_VERSION[0]}.${CURRENT_VERSION[1]}.${CURRENT_VERSION[2]}"
  export NEW_VERSION="${NEW_VERSION[0]}.${NEW_VERSION[1]}.${NEW_VERSION[2]}"
}

create_new_version_file() {
  print "Removing version.go file for v$CURRENT_VERSION"
  rm config/version.go

  print "Creating version.go file at config/version.go for v$NEW_VERSION"

  echo "package config

const VERSION = \"$NEW_VERSION\"" >> config/version.go
}

publish_new_version() {
  set_new_version "$1"
  create_new_version_file

  print "Pushing new version to Github"

  git checkout master
  git add ./config/version.go
  git commit -m "bump version to v$NEW_VERSION"
  git tag v$NEW_VERSION
  git push origin v$NEW_VERSION
  git push

  print "Version $NEW_VERSION successfully published"
}

print() {
    echo "---  $1"
}
