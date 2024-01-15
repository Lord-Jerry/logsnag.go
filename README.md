<div align="center">
	<img src="https://logsnag.com/og-image.png" alt="LogSnag"/>
	<br>
    <h1>LogSnag</h1>
	<p>Get notifications and track your project events.</p>
	<a href="https://discord.gg/dY3pRxgWua"><img src="https://img.shields.io/discord/922560704454750245?color=%237289DA&label=Discord" alt="Discord"></a>
	<a href="https://docs.logsnag.com"><img src="https://img.shields.io/badge/Docs-LogSnag" alt="Documentation"></a>
	<br>
	<br>
</div>


## Installation

```sh

```

## Usage

### Import Library

```go
import LogSnag "logsnag" 
```

### Initialize Client

```go
package main

import LogSnag "logsnag"

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
}
```

### Publish Event

```go
package main

import (
	"fmt"
	LogSnag "logsnag"
)

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
	err := logsnag.Track(&LogSnag.TrackData{
		UserId: "user_123",
		Channel: "waitlist",
		Event: "User Joined",
		Description: "Email: john@doe.com",
		Icon: "🎉",
		Notify: true,
		Tags:  &map[string]any{
			"source": "google",
        },
    })

	if err != nil {
		fmt.Println(err)
	}
}
```

### User Properties

```go
package main

import (
"fmt"
LogSnag "logsnag"
)

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
	err := logsnag.Identify(&LogSnag.IdentifyData{
		UserId: "user_123",
		Properties: &map[string]any{
			"name":  "John Doe",
			"email": "john@doe.com",
			"plan":  "free",
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}
}
```

### Group Properties

```go
package main

import (
	"fmt"
	LogSnag "logsnag"
)

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
	err := logsnag.Group(&LogSnag.GroupData{
		UserId:  "user_123",
		GroupId: "group_123",
		Properties: &map[string]any{
			"name": "ACME Inc.",
			"plan": "enterprise",
			"fund": "Series A",
		},
	})

	if err != nil {
		fmt.Println(err)
	}
}

```

### Publish Insight

```go
package main

import (
	"fmt"
	LogSnag "logsnag"
)

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
	err := logsnag.Insight.Track(&LogSnag.InsightData{
		Title: "User Count",
		Value: 100,
		Icon:  "👨",
	})

	if err != nil {
		fmt.Println(err)
	}
}

```

### Increment Insight

```go
package main

import (
	"fmt"
	LogSnag "logsnag"
)

func main() {
	logsnag := LogSnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas", false)
	err := logsnag.Insight.Increment(&LogSnag.InsightData{
		Title: "User Count",
		Value: 100,
		Icon:  "👨",
	})

	if err != nil {
		fmt.Println(err)
	}
}

```