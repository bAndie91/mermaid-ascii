package cmd

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func verifyMap(t *testing.T, mermaid string, expectedMap string) {
	mermaidMap, styleClasses, err := mermaidFileToMap(mermaid)
	if err != nil {
		log.Fatal("Failed to parse mermaid: ", err)
	}
	actualMap := drawMap(mermaidMap, *styleClasses)
	if expectedMap != actualMap {
		t.Errorf("Map didn't match actual map\nExpected: %v\nActual: %v", expectedMap, actualMap)
	}
}

func TestTwoNodes(t *testing.T) {
	mermaid :=
		`graph LR
A --> B`
	expectedMap :=
		`+---+     +---+
|   |     |   |
| A |---->| B |
|   |     |   |
+---+     +---+`
	verifyMap(t, mermaid, expectedMap)
}

func TestTwoNodesLongerNames(t *testing.T) {
	mermaid :=
		`graph LR
ABC --> BCDEFG`
	expectedMap :=
		`+-----+     +--------+
|     |     |        |
| ABC |---->| BCDEFG |
|     |     |        |
+-----+     +--------+`
	verifyMap(t, mermaid, expectedMap)
}

func TestTwoLayerSingleGraph(t *testing.T) {
	mermaid :=
		`graph LR
A --> B
A --> C`
	expectedMap :=
		`+---+     +---+
|   |     |   |
| A |---->| B |
|   |     |   |
+---+     +---+
  |            
  |            
  |            
  |            
  |            
  |       +---+
  |       |   |
  ------->| C |
          |   |
          +---+`
	verifyMap(t, mermaid, expectedMap)
}

func TestTwoLayerSingleGraphLongerNames(t *testing.T) {
	mermaid :=
		`graph LR
ABC --> BCDEFG
ABC --> CDEFGHI`
	expectedMap :=
		`+-----+     +---------+
|     |     |         |
| ABC |---->|  BCDEFG |
|     |     |         |
+-----+     +---------+
   |                   
   |                   
   |                   
   |                   
   |                   
   |        +---------+
   |        |         |
   -------->| CDEFGHI |
            |         |
            +---------+`
	verifyMap(t, mermaid, expectedMap)
}

func TestThreeNodes(t *testing.T) {
	mermaid :=
		`graph LR
A --> B
B --> C`
	expectedMap :=
		`+---+     +---+     +---+
|   |     |   |     |   |
| A |---->| B |---->| C |
|   |     |   |     |   |
+---+     +---+     +---+`
	verifyMap(t, mermaid, expectedMap)
}

func TestTwoRootNodes(t *testing.T) {
	mermaid :=
		`graph LR
A --> B
C --> D`
	expectedMap :=
		`+---+     +---+
|   |     |   |
| A |---->| B |
|   |     |   |
+---+     +---+
               
               
               
               
               
+---+     +---+
|   |     |   |
| C |---->| D |
|   |     |   |
+---+     +---+`
	verifyMap(t, mermaid, expectedMap)
}

func TestTwoRootNodesLongerNames(t *testing.T) {
	mermaid :=
		`graph LR
ABC --> BCDEFG
CDEFGH --> DEF`
	expectedMap :=
		`+--------+     +--------+
|        |     |        |
|  ABC   |---->| BCDEFG |
|        |     |        |
+--------+     +--------+
                         
                         
                         
                         
                         
+--------+     +--------+
|        |     |        |
| CDEFGH |---->|  DEF   |
|        |     |        |
+--------+     +--------+`
	verifyMap(t, mermaid, expectedMap)
}