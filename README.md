<p align="center">
  <img src="https://github.com/aleixfortm/gowol/assets/95043218/1bce0a50-3bdb-42ef-82c5-ce158e18938f" alt="storymous_tree" width="200" height="auto">
</p>



### **<p align="center">WOLed</p>**


<p align="center"><strong>> An over-engineered CLI Wake-on-LAN tool</strong><br>Add your own devices, edit config data and broadcast WOL packages through a CLI application</p>

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" /
  </a>
</div>


<br>

## Usage
```go
woled <command>
```

## Commands and flags
```go
  wol         Broadcast WOL packet to local network
  config      Show and tweak default app settings, including UDP Port and broadcast IP address
  list        Display a list of saved devices
  save        Save device data to local storage
  remove      Remove saved device
```
```go
  -h, --help   show arguments and examples
```

<br>

## <code>wol</code> [name]
### <strong>Description</strong>
Send a <code>WOL packet</code> to the local network, broadcasted to <code>IP 255.255.255.255</code> by default.

### <strong>Usage</strong>
```python
  woled wol [name]
```
### <strong>Arguments</strong>
```go
  [name]  string  Name of the device
```
### <strong>Examples</strong>
```go
  woled wol PC-1
  woled wol "My computer"
```

<br>

## <code>list</code>
### <strong>Description</strong>
Display a list of previously saved devices from local storage file <code>data.json</code>

### <strong>Usage</strong>
```python
  woled list
```
### <strong>Arguments</strong>
```go
  None
```
### <strong>Examples</strong>
```go
  woled list
```
### <strong>Output</strong>
```
  > Device list:
  0 Main-computer
  1 PC-2
```

<br>

## <code>save</code> [name] [MAC]
### <strong>Description</strong>
Save your device to a local data file by specifying a <code>name</code> and the <code>MAC address</code> of the device.

### <strong>Usage</strong>
```python
  woled save [name] [MAC]
```
### <strong>Arguments</strong>
```go
  [name]  string   Name of the device
  [MAC]   string   MAC address of the device
```
### <strong>Examples</strong>
```go
  woled save Main-computer 22:F4:63:90:A3:75
  woled save "My computer" 04:AF:32:4B:4C:95
```
### <strong>Output</strong>
```
  > Main-computer saved successfully with MAC address 22:F4:63:90:A3:75
```

<br>

## <code>remove</code> [name]
### <strong>Description</strong>
Remove saved device from local storage data <code>data.json</code>

### <strong>Usage</strong>
```python
  woled remove [name]
```
### <strong>Arguments</strong>
```go
  [name]  string  Name of the device
```
### <strong>Examples</strong>
```go
  woled remove PC-1
  woled remove "My computer"
```

### <strong>Output</strong>
```
  > Main-computer has been successfully deleted.
```

<br>


# Data.json file structure
## Model
```Python
Device:
  Name:        string
  MACAddress:  string
```

<br>

## Example:
```js
[
    {
        "name": "Main-computer",
        "macAddress": "22:F4:63:90:A3:75"
    },
    {
        "name": "PC-2",
        "macAddress": "5E:F9:AA:70:A3:8C"
    }
]
```
