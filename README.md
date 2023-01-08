# d7

Lightweight AUR helper written in Golang.

### Features

- Add packages.
- Update existing packages.
- Clean your cloned packages directory.
- View a package's PKGBUILD.

### Why?

I just wanted a minimal AUR helper. The code is ugly but works. Also my first project in Golang.

### Compiling

```
go install
go build .
```

### Usage

#### Add

```
d7 add <package_name>
```

#### Sync

```
d7 sync
```

#### Clean

```
d7 clean
```

#### View PKGBUILD

```
d7 view <pkg_name>
```
