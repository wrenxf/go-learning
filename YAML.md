在后端开发中，数据序列化（将内存中的数据结构转换为可存储或网络传输的字节流）是极其核心的基础能力。除了你之前了解的 JSON 和 YAML，业界还有多种序列化格式。

下面梳理了 5 种最常用的序列化格式，并从底层原理、优缺点及适用场景进行了全面对比：

### **1. JSON (JavaScript Object Notation)**

- **特点**：纯文本格式，人类可读性极强，自带键值对结构。
- **优点**：跨语言支持最完善，生态极其繁荣；调试极其方便（直接看报文即可）。
- **缺点**：文本格式导致体积较大；解析需要大量字符串处理，性能相对较低；不支持二进制数据（如图片、音频）的直接传输。
- **适用场景**：Web 前后端 API 通信（RESTful）、微服务间的普通 RPC 调用、NoSQL 文档存储（如 MongoDB）。

### **2. Protocol Buffers (Protobuf)**

- **特点**：由 Google 开源的二进制序列化格式，需要预先定义 `.proto` 接口文件。
- **优点**：**性能怪兽**。二进制格式使得序列化后的体积比 JSON 小 3~10 倍，解析速度比 JSON 快 20~100 倍；自带强类型约束和代码生成工具，能极大减少 Bug。
- **缺点**：人类无法直接阅读生成的二进制数据，调试相对困难；每次修改数据结构都需要重新生成代码。
- **适用场景**：微服务内部的高并发 RPC 通信（如 gRPC 的默认协议）、对带宽和延迟要求极高的内部系统。

### **3. YAML (YAML Ain't Markup Language)**

- **特点**：专为人类阅读设计的文本格式，使用缩进表示层级，支持注释。
- **优点**：可读性在所有格式中最好，非常适合表达复杂的嵌套配置。
- **缺点**：解析器逻辑复杂，性能较差；对缩进（空格）极其敏感，少一个空格就会导致解析失败。
- **适用场景**：项目配置文件（如 Spring Boot、K8s、Docker Compose、CI/CD 流水线）。**绝不用于网络数据传输**。

### **4. MessagePack (Msgpack)**

- **特点**：被称为“二进制版的 JSON”。它的 API 和数据结构与 JSON 几乎一模一样，但底层将数据压缩成了二进制。
- **优点**：兼具了 JSON 的易用性和二进制的高性能。体积比 JSON 小，解析比 JSON 快；不需要像 Protobuf 那样提前写 `.proto` 文件，极其灵活。
- **缺点**：由于是二进制格式，无法像 JSON 那样直接肉眼阅读报文。
- **适用场景**：Redis 缓存存储（节省内存）、需要高性能但又不想维护 `.proto` 文件的微服务通信、游戏后端服务。

### **5. XML (eXtensible Markup Language)**

- **特点**：老牌的标签式文本格式，支持极其严格的 Schema 校验。
- **优点**：扩展性极强，自描述性好，支持复杂的命名空间和类型校验。
- **缺点**：标签冗余导致体积极其庞大；解析极其耗费 CPU 和内存。
- **适用场景**：传统的企业级遗留系统（如银行、金融行业的 SOAP 接口）、某些特定的硬件设备通信。在新项目中已极少使用。

------

### **💡 核心对比与选型总结**

表格

| 格式            | 类型   | 体积 | 性能 | 可读性 | 核心定位               |
| :-------------- | :----- | :--- | :--- | :----- | :--------------------- |
| **JSON**        | 文本   | 大   | 较慢 | 极佳   | 对外 API、通用数据交换 |
| **Protobuf**    | 二进制 | 极小 | 极快 | 差     | 内部高并发 RPC、gRPC   |
| **YAML**        | 文本   | 中等 | 慢   | 极佳   | 配置文件、基础设施     |
| **MessagePack** | 二进制 | 较小 | 快   | 差     | 缓存、灵活的高性能传输 |
| **XML**         | 文本   | 极大 | 极慢 | 一般   | 传统遗留系统、SOAP     |

**Go 后端开发建议：**

- **对外提供接口**：无脑选 **JSON**（Gin 的 `c.JSON()`）。
- **内部微服务通信**：强烈建议学习并使用 **Protobuf + gRPC**，这是目前大厂的标准配置。
- **写项目配置**：使用 **YAML**（配合 `viper` 库）。
- **存入 Redis**：如果数据量大，考虑使用 **MessagePack** 替代 JSON，能为你省下大量服务器内存。



# YAML 使用笔记（Go 后端视角）

> 面向 Go 后端开发者的 YAML 学习笔记：先讲**为什么**（定位与设计动机），再讲**语法**（核心机制），最后讲**在 Go 里怎么用**（yaml.v3 + 常见坑）。
> 适用场景：配置文件（config.yaml）、Docker Compose、K8s 清单、CI/CD、OpenAPI 定义。

---

## 1. YAML 是什么：先搞清楚它的定位

**YAML = YAML Ain't Markup Language**（递归缩写，"YAML 不是标记语言"）。

一句话定位：**YAML 是一种"以人类可读为目标"的数据序列化格式，专门为配置文件而生。**

它和 JSON 同源——**YAML 1.2 是 JSON 的超集**（JSON 文档本身就是合法的 YAML 文档）。所以你可以粗暴理解：

> JSON  = 机器友好、人类勉强可读的数据格式
> YAML  = 人类友好、机器可解析的配置格式

### 1.1 三大数据格式对比（先建立坐标系）

| 维度 | JSON | YAML | TOML |
|---|---|---|---|
| 设计目标 | 数据交换（跨语言） | 配置文件（人类可读） | 配置文件（简单、无歧义） |
| 缩进语法 | 无（全靠 `{}` `[]`） | **空格缩进**（核心） | 无（靠 `[section]` 分组） |
| 注释支持 | ❌ 不支持 | ✅ `#` | ✅ `#` |
| 多行字符串 | 麻烦 | ✅ `\|` `>` 原生支持 | 一般 |
| 类型推断 | 显式类型 | 自动推断（有陷阱） | 显式/自动混合 |
| 复杂度 | 低 | **高**（语法特性多） | 低 |
| 学习成本 | 低 | 中（缩进 + 类型陷阱） | 低 |

**什么时候用 YAML**：写配置文件、CI/CD（GitHub Actions、GitLab CI）、容器编排（Docker Compose、K8s）、OpenAPI/Swagger。
**什么时候别用 YAML**：程序间数据交换（用 JSON/Protobuf）、需要严格类型和低歧义的场景（考虑 TOML）。

> 机理要点：YAML 之所以"可读"，是因为它把**结构信息（层级、分组）交给缩进和换行**来表达，而不是交给 `{ } [ ] ,` 这些符号。牺牲了一点严格性，换来了极佳的观感。代价就是——**缩进错了，解析就错**。

---

## 2. 核心语法（按"构造块"逐个击破）

### 2.1 最小单元：键值对（Key-Value）

```yaml
name: tiny-feed        # 键: 空格 + 值
port: 8080
```

**铁律：冒号 `:` 后面必须跟一个空格**（否则它只是字符串的一部分）。

```yaml
name:tiny-feed   # ❌ 解析不出来，整个是字符串 "name:tiny-feed"
name: tiny-feed  # ✅
```

### 2.2 层级：靠缩进表达（YAML 的灵魂）

```yaml
server:
  host: 0.0.0.0
  port: 8080
database:
  host: localhost
  port: 3306
```

三条规则：
1. **用空格缩进，绝对不能用 Tab**（Tab 直接报错）。
2. 同级元素**缩进必须完全一致**（对齐）。
3. 缩进量本身多大无所谓（2 格、4 格都行），**关键是同层一致**。

```yaml
server:
  host: 0.0.0.0
    port: 8080   # ❌ port 缩进比 host 多，层级错乱 → 解析报错
```

### 2.3 列表（数组）：`- ` 开头

```yaml
# 顶层列表
- mysql
- redis
- kafka

# 列表里放对象（最常见的写法）
servers:
  - name: backend
    port: 8080
  - name: frontend
    port: 3000
```

机理：`- item` 中的 `-` 和 `item` 之间要有空格；列表项内部的键继续按缩进对齐。

### 2.4 数据类型与自动类型推断（重要陷阱区）

YAML 会**根据字面量自动推断类型**。这是便利，也是坑源。

```yaml
int: 8080              # 整数
float: 3.14            # 浮点
bool_true: true        # 布尔
bool_false: false      # 布尔
null1: null            # 空值
null2: ~               # 空值（另一种写法）
date: 2026-08-31       # 日期（ISO 8601）→ 解析为 time.Time
str1: hello            # 字符串
str2: "8080"           # 引号强制字符串
str3: 'hello world'    # 单引号：原样字符串
```

**三个高频坑（必须记住）：**

| 写法 | 推断结果 | 后果 |
|---|---|---|
| `port: 8080` | int | Go 里 unmarshal 到 `string` 字段会**报错** |
| `port: "8080"` | string | 安全，但注意是字符串 |
| `version: 1.0` | float（不是 string） | unmarshal 到 `string` 会报错 |
| `enable: yes` | **yaml.v2 里是 bool，v3 里是 string** | 版本不同行为不同（见 §4） |

> 机理要点：YAML 1.1 规范里 `yes/no/on/off` 会被当作布尔值；YAML 1.2（及 Go 的 `yaml.v3`）**收紧了规则，只有 `true/false` 是布尔**。这就是为什么老代码里 `enable: yes` 能跑、新代码行为却变了——**库版本决定类型语义**。

### 2.5 字符串的引号规则

| 写法 | 含义 |
|---|---|
| `key: hello` | 裸字符串，简单值可用 |
| `key: "hello world"` | 双引号：支持 `\n` `\t` 转义 |
| `key: 'hello world'` | 单引号：不转义，原样输出 |
| `key: "a: b"` | 值里含冒号+空格 → **必须加引号** |
| `key: "#comment"` | 值以 `#` 开头 → 必须加引号 |
| `key: "yes"` | 想让它保持字符串 → 加引号 |

判断法则：**如果值里有 `: `（冒号空格）、`#`、`- `、特殊开头字符，或你想强制字符串类型，就加引号。** 不确定时加引号永远安全。

### 2.6 多行字符串：块标量 `|` 和 `>`

配置里写 SQL、脚本、长文本时的神器：

```yaml
# | 保留换行（literal）
sql: |
  SELECT *
  FROM users
  WHERE id = ?
  LIMIT 10;

# > 折叠换行（folding）：连续行合并成一行，空行才产生换行
desc: >
  this is a long
  description that will
  be folded into one line

# 加上 - 去掉末尾换行符；加上 + 保留所有末尾换行
no_trailing: |-
  line1
  line2
keep_all: |+
  line1
  line2

```

| 符号 | 行为 | 末尾换行 |
|---|---|---|
| `\|` | 保留内部所有换行 | 保留一个 |
| `>-` | 折叠换行 | 去掉 |
| `\|-` | 保留换行 | 去掉 |
| `\|+` | 保留换行 | 保留全部（含多个空行） |

### 2.7 锚点 & 别名（`&` 定义，`*` 引用）—— 去重利器

配置文件里大量重复结构（比如多个环境、多个副本）时用：

```yaml
defaults: &defaults
  timeout: 30
  retries: 3

dev:
  <<: *defaults          # 合并键：继承 defaults 的所有字段
  url: http://dev.local

prod:
  <<: *defaults
  url: http://prod.example.com
  timeout: 60            # 覆盖 defaults 里的 timeout
```

- `&defaults`：给这个节点起名（锚点）。
- `*defaults`：引用这个节点（别名）。
- `<<:`：合并键，把锚点里的键展开进当前映射。

**注意：锚点只是 YAML 层面的语法糖，Go 的 `yaml.v3` 完全支持，但合并后的值在解码时才展开。**

---

## 3. Go 后端实战：yaml.v3

### 3.1 库的选择：`gopkg.in/yaml.v3`（v2 已过时）

```bash
go get gopkg.in/yaml.v3
```

**为什么用 v3 不用 v2：**

| 对比项 | yaml.v2 | yaml.v3 |
|---|---|---|
| 规范 | YAML 1.1（`yes/no` 当布尔） | YAML 1.2（只有 `true/false`） |
| 解析器 | 老实现 | 重写，性能更好 |
| Node API | ❌ | ✅ `yaml.Node` 可编程操作 |
| 严格模式 | 无 | ✅ `Decoder.KnownFields(true)` |

### 3.2 标准用法：Unmarshal 到 struct

```go
package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// struct tag 与 JSON 完全同构：yaml:"字段名"
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Debug    bool           `yaml:"debug"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func main() {
	// 1. 读文件
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	// 2. 反序列化（unmarshal）
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", cfg)
	// 输出: {Server:{Host:0.0.0.0 Port:8080} Database:{Host:localhost Port:3306} Debug:false}
}
```

对应 `config.yaml`：

```yaml
server:
  host: 0.0.0.0
  port: 8080
database:
  host: localhost
  port: 3306
  user: root
  password: secret
  dbname: tinyfeed
debug: false
```

**关键机制：struct tag 的匹配规则**
- `yaml:"server"` 决定 YAML 里的键名。
- 大小写不敏感匹配：yaml 里的 `Server` 也能匹配 tag 为 `server` 的字段。
- **字段名首字母必须大写**（导出），否则库看不到这个字段。

### 3.3 struct tag 的常用选项

```go
type Config struct {
	Server   ServerConfig `yaml:"server"`
	Port     int          `yaml:"port,omitempty"` // 值为零值(0)时不输出
	Tags     []string     `yaml:"tags,flow"`      // 输出为 JSON 风格 [a, b]
	Internal string       `yaml:"-"`              // 完全忽略该字段
	Dynamic  *DynamicCfg  `yaml:"dynamic,inline"` // 内联展开字段（不嵌套）
}
```

| 选项 | 作用 |
|---|---|
| `omitempty` | 序列化时，零值字段不输出（Marshal 时有用） |
| `flow` | 列表/映射输出成 `[a, b]` 行内格式 |
| `-` | 忽略该字段 |
| `,inline` | 结构体字段内联，不产生嵌套层级 |

### 3.4 序列化（Marshal）：把 struct 写回 YAML

```go
cfg := Config{
	Server: ServerConfig{Host: "0.0.0.0", Port: 8080},
	Debug:  true,
}

out, err := yaml.Marshal(&cfg)
if err != nil {
	panic(err)
}
fmt.Println(string(out))
```

输出：

```yaml
server:
    host: 0.0.0.0
    port: 8080
debug: true
```

### 3.5 生产级读取流程：Decoder + KnownFields（强烈推荐）

只用 `yaml.Unmarshal` 有个隐患：**yaml 里有键但 struct 没对应字段时，默认静默忽略**——配置写错名字，程序不报错，行为诡异。用 `Decoder` 开启严格模式：

```go
f, err := os.Open("config.yaml")
if err != nil {
	panic(err)
}
defer f.Close()

dec := yaml.NewDecoder(f)
dec.KnownFields(true) // 遇到未知字段 → 直接报错！

var cfg Config
if err := dec.Decode(&cfg); err != nil {
	panic(err) // yaml: unmarshal errors: line 3: field serverr not found in type Config
}
```

> 机理：`KnownFields(true)` 让解码器在 YAML 键与 struct 字段**一一对应**，拼错键名立刻暴露。生产配置强烈建议开启。

### 3.6 进阶：yaml.Node —— 不预先定义 struct 也能读

```go
var root yaml.Node
yaml.Unmarshal(data, &root) // root 是文档节点树

// 遍历找 "server" 下的 "port"
for _, doc := range root.Content {
	for _, kv := range doc.Content {
		if kv.Value == "server" {
			// kv.Content[0] 是 key，kv.Content[1] 是 value
			fmt.Println(kv.Content[1].Content[3].Value) // port 的值
		}
	}
}
```

适用场景：配置结构不固定、需要动态读取的配置系统（比如实现自己的配置热加载）。

---

## 4. 常见坑清单（按踩坑频率排序）

### 坑 1：Tab 缩进 / 缩进不一致
```
yaml: line 3: found character that cannot start any token
```
**解法：** 编辑器统一用空格（VS Code 装 YAML 插件，`editor.insertSpaces: true`）；注意复制网上的 YAML 可能带 Tab。

### 坑 2：数字类型不匹配
```go
type C struct {
	Port string `yaml:"port"`
}
// yaml: unmarshal errors: cannot unmarshal !!int `8080` into string
```
**解法：** YAML 侧写 `"8080"`，或者 Go 侧用 `int`。**配置文件里端口、超时这类"看起来像数字的配置"，明确想好用 int 还是 string。**

### 坑 3：`yes/no` 的版本差异
```yaml
enable: yes   # yaml.v2 → 解析为 true；yaml.v3 → 解析为字符串 "yes"
```
**解法：** 统一写 `true/false`，不要用 `yes/no/on/off`。

### 坑 4：未知字段被静默忽略
yaml 里写了 `serverr`（拼错），Unmarshal 不报错，程序用默认值运行 → 线上事故。
**解法：** 见 §3.5，用 `Decoder.KnownFields(true)`。

### 坑 5：时间解析
```yaml
created: 2026-08-31 16:00:00   # ❌ 有空格，不是合法 ISO 8601
created: 2026-08-31T16:00:00Z  # ✅ 带 T 和时间
```
**解法：** YAML 里的时间必须符合 ISO 8601（`T` 分隔），否则 unmarshal 到 `time.Time` 报错。只想当字符串存，就加引号。

### 坑 6：值里包含特殊字符忘记加引号
```yaml
password: p@ss:word    # ❌ 冒号+无空格其实可以，但 p@ss: word 就炸了
url: http://example.com # 无空格冒号可以裸写，但建议统一加引号
```
**解法：** 不确定就加引号。含 `: `、`#`、`{`、`[`、`- ` 开头的值必须引号。

---

## 5. 实战结合：tiny-feed 风格的 config.yaml

一个典型的 Go 后端服务配置骨架（对照你的 tiny-feed 复刻项目）：

```yaml
app:
  name: tiny-feed
  env: dev            # dev | prod
  debug: true
  port: 8080

server:
  read_timeout: 10s
  write_timeout: 30s
  graceful_shutdown: 5s

database:
  driver: mysql
  host: 127.0.0.1
  port: 3306
  user: root
  password: "123456"
  dbname: tinyfeed
  charset: utf8mb4
  max_open_conns: 100
  max_idle_conns: 10

jwt:
  secret: "change-me-in-production"
  expire: 24h

redis:
  host: 127.0.0.1
  port: 6379
  db: 0

log:
  level: info          # debug | info | warn | error
  output: stdout       # stdout | file
```

对应的 Go 结构体骨架（model → 与前面一致，字段一一对应）：

```go
type Config struct {
	App      AppConfig      `yaml:"app"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	JWT      JWTConfig      `yaml:"jwt"`
	Redis    RedisConfig    `yaml:"redis"`
	Log      LogConfig      `yaml:"log"`
}
```

加载入口：

```go
func Load(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	dec.KnownFields(true)

	var cfg Config
	if err := dec.Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
```

---

## 6. 语法速查表（一页背下来）

| 语法 | 写法 | 说明 |
|---|---|---|
| 键值对 | `key: value` | 冒号后**必须空格** |
| 嵌套 | 增加缩进 | 同级对齐，禁 Tab |
| 列表 | `- item` | 短横线后空格 |
| 注释 | `# 注释` | 仅支持行注释 |
| 多行保留换行 | `key: \|` | 块标量 |
| 多行折叠 | `key: >` | 合并为一行 |
| 强制字符串 | `"..."` / `'...'` | 双引号支持转义 |
| 布尔 | `true` / `false` | 别用 yes/no |
| 空值 | `null` / `~` | 两者等价 |
| 锚点 | `&name` | 定义复用节点 |
| 别名 | `*name` | 引用锚点 |
| 合并键 | `<<: *name` | 继承并覆盖 |

