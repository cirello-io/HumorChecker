# HumorChecker - port of SentiMental into Go.

[![Build Status](https://travis-ci.org/ccirello/HumorChecker.svg?branch=master)](https://travis-ci.org/ccirello/HumorChecker)

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

import hc "cirello.io/HumorChecker"

func main(){
	hc.Analyze("Hey you worthless scumbag"); //Score: -6, Comparative:-1.5
	hc.Positivity("This is so cool"); //Score: 1, Comparative:.25
	hc.Negativity("Hey you worthless scumbag"); //Score: 6, Comparative:1.5
	hc.Analyze("I am happy"); //Score: 3, Comparative: 1
	hc.Analyze("I am so happy"); //Score: 6, Comparative: 1.5
	hc.Analyze("I am extremely happy"); //Score: 12, Comparative: 3
	hc.Analyze("I am really sad"); //Score: -4, Comparative: -1
}
```

