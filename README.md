# git
<img src="docs/img/badges.svg">

Git and GitHub handlers for WebTyp projects. Extracted from `webtyp/devflow` so
consumers (e.g. `webtyp/sitepub`) don't drag the whole devflow dependency tree.

## Features

- **Git handler** (`NewGit`): add/commit/tag/push, clone/pull/fetch, tags, `.gitignore` entries.
- **GitHub client** (`NewGitHub`): repos, releases, gh CLI integration.
- **GitHub auth**: OAuth Device Flow (`NewGitHubOAuth`), PAT recovery (`NewGitHubPATAuth`), `EnsureGHSession`.
- **Keyring** (`NewKeyring`): secure token storage via the system keyring.
- **Secrets** (`GitHubSecrets`): repository secrets via gh CLI.
- **TestCache** (`NewTestCache`): git-based test cache used by `gotest`.
- **Helpers**: `CompareVersions`, commit message builders, `WorkTreeDirtyBeyond`, publish objector types (`PublishContext`, `PublishAction`).

## Usage

```go
import "webtyp.com/git"

g, err := git.NewGit()
if err != nil {
    log.Fatal(err)
}
```

## Installing tools

Use `goinstall` to (re)build the WebTyp CLI tools from `webtyp/devflow`.

## License

MIT