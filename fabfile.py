#!/usr/bin/python3
# encoding=utf-8

from fabric.api import *
import os
from  datetime import datetime

env.user='root'
env.password=''

env.hosts=['root@xxx']
env.roledefs={
  "server1":['root@xxx'],
}
env.port=22

EXEFILE='readbook'
EXEPATH='/opt/readbook'
_TAR_FILE='readbook.tar.gz'
DT=datetime.now().strftime('%Y-%m-%d_%H:%M:%S')
LNEXEFILE=('%s.%s' %(EXEFILE,DT))

def build():
  local('rm -f %s*'%EXEFILE)
  version=local('git describe --tags --dirty --always',capture=True)
  local('GOOS=linux  go build -i -o %s -ldflags "-X main.Version=%s"' %(LNEXEFILE,version))

def tar():
  build()
  with settings(warn_only=True):
    includes=['conf','static','%s'%LNEXEFILE]
    unexcludes=['*.go','*.lock','vendor','*.toml','.*','test','*.pyc', '*.pyo']
    
    with lcd(os.path.abspath('.')):
      cmd = ['tar', '--dereference', '-czvf', './dist/%s' % _TAR_FILE]
      cmd.extend(['--exclude=\'%s\'' % ex for ex in unexcludes])
      cmd.extend(includes)
      local(' '.join(cmd))

def firstpush():
    with settings(warn_only=True):
      put('./dist/%s'% _TAR_FILE,EXEPATH)
      with cd(EXEPATH):
        run('tar -xzvf %s'%_TAR_FILE)
        start()

def push():
  build()
  with settings(warn_only=True):
    put(LNEXEFILE,EXEPATH)
    start()

def start():
  with settings(warn_only=True):
    with cd(EXEPATH):
      run('rm -f %s'%EXEFILE)
      run('ln -s %s %s'%(LNEXEFILE,EXEFILE))
      run('chmod 777 %s' %EXEFILE)
      run('supervisorctl stop %s' %EXEFILE)
      run('supervisorctl start %s' %EXEFILE)

def restart():
  with settings(warn_only=True):
    with cd(EXEPATH):
      run('supervisorctl stop %s' %EXEFILE)
      run('supervisorctl start %s' %EXEFILE)