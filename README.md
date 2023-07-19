# Jira

Parse Jira issues from text.

## Installation

```bash
go install github.com/mskelton/jira@latest
```

## Prerequisites

Add the following environment variables to your shell with your Jira base URL
and supported prefixes.

```bash
export JIRA_BASE_URL=https://my.atlassian.net
export JIRA_PREFIXES="ABC,SUPPORT"
```

## Usage

**Input**

```bash
echo "Check out SUPPORT-123 on Jira" | jira
```

**Output**

```
https://my.atlassian.net/browse/SUPPORT-123
```
