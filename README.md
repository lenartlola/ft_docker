# FT_DOCKER

**Trying to make my own docker in golang.**

## Usage:
```sh
sudo go run main.go run /bin/sh
```

## Dependencies:
- vagrant
```shell
pacman -Syu vagrant
```

## Namespaces:
Namespaces are one of a feature in the Linux Kernel and fundamental aspect of containers on Linux. On the other hand, namespaces provide a layer of isolation. 

From wikipedia:
> Namespaces are a feature of the Linux kernel that partitions kernel
> resources such that one set of processes sees one set of resources and
> another set of processes sees a different set of resources. The feature
> works by having the same namespace for a group of resources and
> processes, but those namespaces refer to distinct resources.


### Namespace Types:
1) Process ID or PID: When a process is created on most Unix-like operating systems, it is given a specific numeric identifier called PID. PID helps to identify a process uniquely even if there are two process that share the same human-readable name.
2) Net namespace: every computer connected to a network requires an IP address. This is unique number that allows computers to communicate effectively.
3) UTS or Unix Timesharing System: it allows for the segregation of hostnames.
4) User namespace: Every computer system has some way of tracking which user owns which file(s). This allows the system to restrict access to sensitive system files. It also prevents people using the same computer from accessing each others' files.
5) MNT or Mount: The mount namespace is used to isolate mount points such that processes in different namespaces cannot view each others' files.
6) IPC or Interprocess communication: [Refer to this](https://opensource.com/article/19/4/interprocess-communication-linux-storage)

Container technologies place user's process into all of the above namespaces and init the namespaces to provide a standard environments.

## CGroups:
CGroups collect a set of process or task ids together and apply limits to them.
CGroups are exposed by the kernel as a special file system you can mount. You add a process or thread to a cgroup by simply adding process ids to a tasks file, and then read and configure various values by essentially editing files in that directory.

## Layered Filesystems:
Namespaces and CGroups are the isolation and resource sharing sides of containerisation. They’re the big metal sides and the security guard at the dock. Layered Filesystems are how we can efficiently move whole machine images around: they're why the ship floats instead of sinks.

