## About Alerts

An alert involves a human. It may be in the middle of the night. Humans can only react to one thing at a time. It must be used sparingly. Prometheus.io has [a good primer on alerting](https://prometheus.io/docs/practices/alerting/). The Google SRE book has a [great chapter on monitoring distributed systems](https://landing.google.com/sre/book/chapters/monitoring-distributed-systems.html).

A well-understood alert has these components:

- At least one metric that can measure it
- A query that can be used to view alerting state
- A way to reproduce alerting state
- An alert rule
- A documented debug path that answers, "What do I do when I see this alert?" This may include related metrics or logging queries, pre-defined graphs, standard operating procedure (SOP) documentation and/or a link to a knowledge base.

Ideally, reproducing the alerting state can be automated to aid development and automated testing.

