<p align="center">
  <img src="https://github.com/aleixfortm/gowol/assets/95043218/1bce0a50-3bdb-42ef-82c5-ce158e18938f" alt="storymous_tree" width="200" height="auto">
</p>



### **<p align="center">Woled</p>**


<p align="center"><strong>> An over-engineered CLI Wake-on-LAN tool</strong><br>Add your own devices, edit config data and broadcast WOL packets through a CLI application</p>

<div align="center">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" /
  </a>
</div>


<br>

## Features
* 🌐 Send WOL packets to your local network
* ✅ User-friendliness and simplicity of a CLI tool
* 🔍 Save, remove and list your own devices


## Getting started 
### Dependencies
| Dependency       | Purpose                 | Version                    | Download link         |
|------------------|-------------------------|----------------------------| ----------------------|
| Go               |  Compiling              | v1.20.6                    | https://go.dev/dl/    |

### Installing and using the app
To install the <code>CLI</code> application in your machine, copy the following code snippet into your terminal:

```cmd
git clone https://github.com/aleixfortm/woled
cd woled
go install
```
You should now be able to use Woled commands from your <code>CLI</code>.

## Usage
```go
woled <command>
```

### Commands
```go
  wol         Broadcast WOL packet to local network
  add         Add device data to local storage
  remove      Remove a device from local storage
  list        Display a list of added devices
  config      Show and tweak default app settings, including UDP Port and broadcast IP address
```

### Flags
```go
  -h, --help   show arguments and examples
```

<br>
<!--
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
Display a list of previously added devices from local storage file <code>data.json</code>

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

## <code>add</code> [name] [MAC]
### <strong>Description</strong>
add your device to a local data file by specifying a <code>name</code> and the <code>MAC address</code> of the device.

### <strong>Usage</strong>
```python
  woled add [name] [MAC]
```
### <strong>Arguments</strong>
```go
  [name]  string   Name of the device
  [MAC]   string   MAC address of the device
```
### <strong>Examples</strong>
```go
  woled add Main-computer 22:F4:63:90:A3:75
  woled add "My computer" 04:AF:32:4B:4C:95
```
### <strong>Output</strong>
```
  > Main-computer added successfully with MAC address 22:F4:63:90:A3:75
```

<br>

## <code>remove</code> [name]
### <strong>Description</strong>
Remove added device from local storage data <code>data.json</code>

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


## Data file structure
### Model
```Python
Device:
  Name:        string
  MACAddress:  string
```

<br>

### Example
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
-->
