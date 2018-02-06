#!/usr/bin/python3
# encoding=utf-8

from fabric.api import *
import os

env.user='root'
env.password=''

env.hosts=['root@']
env.roledefs={
  "server1":['root@'],
}
env.port=22

EXEFILE='readbook'
EXEPATH='/opt/readbook'
_TAR_FILE='readbook.tar.gz'

def  a():
  a=local('git describe --tags --dirty --always',capture=True)
  print(a)

def build():

  local('GOOS=linux  go build -i -o $app -ldflags "-X main.Version=${VERSION}"')

  with settings(warn_only=True):
    includes=['conf','static',EXEFILE]
    unexcludes=['*.go','*.lock','vendor','*.toml','.*','test','*.pyc', '*.pyo']
    
    with lcd(os.path.abspath('.')):
      cmd = ['tar', '--dereference', '-czvf', './dist/%s' % _TAR_FILE]
      cmd.extend(['--exclude=\'%s\'' % ex for ex in unexcludes])
      cmd.extend(includes)
      local(' '.join(cmd))

def run():
  build()
  with settings(warn_only=True):
    put(EXEFILE,EXEPATH)
    with cd(EXEPATH):
      run('chmod 777 %s' %EXEFILE)
    