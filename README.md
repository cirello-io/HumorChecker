# HumorChecker - port of SentiMental into Go.

[![Build Status](https://travis-ci.org/ucirello/HumorChecker.svg?branch=master)](https://travis-ci.org/ucirello/HumorChecker)

Credits where credits are due: consider starring [SentiMental](https://github.com/thinkroth/Sentimental) repo on which this one was based.

Sentiment analysis tool based on the [AFINN-111 wordlist](http://www2.imm.dtu.dk/pubdb/views/publication_details.php?id=6010).

## Install
    $ go get cirello.io/HumorChecker

## Features

  * Positivity ranking
  * Negativity ranking
  * Analyze - combines Positivity and Negativity ranking into an aggregate sentiment score

## Example
```js
package main

import (
	"fmt"

	hc "cirello.io/HumorChecker"
)

func main() {
	fmt.Printf("%#v\n", hc.Analyze("Hey you worthless scumbag"))
	fmt.Printf("%#v\n", hc.Positivity("This is so cool"))
	fmt.Printf("%#v\n", hc.Negativity("Hey you worthless scumbag"))
	fmt.Printf("%#v\n", hc.Analyze("I am happy"))
	fmt.Printf("%#v\n", hc.Analyze("I am so happy"))
	fmt.Printf("%#v\n", hc.Analyze("I am extremely happy"))
	fmt.Printf("%#v\n", hc.Analyze("I am really sad"))
}

```

