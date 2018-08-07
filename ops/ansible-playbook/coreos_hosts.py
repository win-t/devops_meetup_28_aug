#!/usr/bin/env python

from __future__ import print_function

import sys
import os
import re
import json

group={
  "lb": ["vm1"],
  "registry": ["vm1"],
  "database": ["vm2"],
  "web": ["vm3", "vm4", "vm5"],
}

os.chdir(os.path.dirname(__file__))

def list_all():
  all = dict()
  for vm_id in range(1,100):
    if os.path.isdir("../coreos-vms/vms/%d/.lock" % vm_id):
      all["vm%d" % vm_id] = { "hosts": ["10.12.34.%d" % (100+vm_id)] }
  for k, vs in group.items():
    all[k] = { "hosts": [] }
    for v in vs:
      g = all.get(v)
      if not g:
        continue
      all[k]["hosts"].extend(g["hosts"])
      all[k]["hosts"] = list(set(all[k]["hosts"]))
  print(json.dumps(all, indent=2))

def show_host(host):
  pattern = re.compile(r'^10\.12\.34\.(\d{1,3})$')
  l = pattern.findall(host)
  def host_not_found():
    print("host not found", file=sys.stderr)
    exit(1)
  if len(l) != 1:
    host_not_found()
  vm_id=int(l[0])-100
  if not os.path.isdir("../coreos-vms/vms/%d/.lock" % vm_id):
    host_not_found()
  print(json.dumps({
    "ansible_user": "core"
  }, indent=2))

if len(sys.argv) == 2 and sys.argv[1] == "--list":
  list_all()
elif len(sys.argv) == 3 and sys.argv[1] == "--host":
  show_host(sys.argv[2])
else:
  print("Usage: %s --list" % sys.argv[0], file=sys.stderr)
  print("Usage: %s --host <host>" % sys.argv[0], file=sys.stderr)
  exit(1)
