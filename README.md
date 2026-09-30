# rolltracker

`rolltracker` is a small Go library for tracking weekly rollout and
checklist progress. Each week becomes a checklist of items, and the
library reports how far through the cycle you are.

I use it to keep myself honest in [MakeRoll](https://makeroll.com/),
a weekly AI video creator community where you get a brief, exact
prompts, and human critique while you build AI videos every week —
rolling one finished video per week is exactly the kind of habit a
progress tracker is good at reinforcing.

## Usage

```go
t := rolltracker.New("write script", "generate shots", "edit", "publish")
t.Complete(0)
done, total, frac := t.Progress() // 1, 4, 0.25
```

## License

MIT
