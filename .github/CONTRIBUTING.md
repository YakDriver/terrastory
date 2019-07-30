# Contributing to Terrastory

**First:** if you're unsure or afraid of _anything_, just ask or submit the issue or pull request anyways. You won't be yelled at for giving your best effort. The worst that can happen is that you'll be politely asked to change something. We appreciate any sort of contributions, and don't want a wall of rules to get in the way of that.

Contributions are welcome, and they are greatly appreciated! Every little bit helps, and credit will always be given.

## Bug Reports

When [reporting a bug][0] please include:

*   Your operating system name and version.
*   Any details about your local setup that might be helpful in troubleshooting.
*   Detailed steps to reproduce the bug.

## Documentation Improvements

Terrastory could always use more documentation, whether as part of the official
Terrastory docs, in docstrings, or even on the web in blog posts, articles, and
such. The official documentation is maintained within this project in
docstrings or in the [docs][3] directory. Contributions are
welcome, and are made the same way as any other code.

## Feature Requests and Feedback

The best way to send feedback is to [file an issue][0] on GitHub.

If you are proposing a feature:

*   Explain in detail how it would work.
*   Keep the scope as narrow as possible, to make it easier to implement.
*   Remember that this is a community-driven, open-source project, and that
    code contributions are welcome. :)

## Development Guide

To set up Terrastory for local development:

1.  Fork [Terrastory](https://github.com/plus3it/terrastory) (look for the
    "Fork" button).

2.  Clone your fork locally:

    ```shell
    git clone https://github.com/your_name_here/terrastory.git && cd terrastory
    ```

3.  Create a branch for local development:

    ```shell
    git checkout -b name-of-your-bugfix-or-feature
    ```

4.  Now you can make your changes locally.

5.  When you're done making changes, use Make to run the linters, the tests,
    and the doc builder. (WIP)

6.  Commit your changes and push your branch to GitHub:

    ```shell
    git add .
    git commit -m "Your detailed description of your changes."
    git push origin name-of-your-bugfix-or-feature
    ```

7.  Submit a pull request through the GitHub website.

## Pull Request Guidelines

If you need some code review or feedback while you are developing the code just
open the pull request.

For pull request acceptance, you should:

1.  Make sure CI (e.g., Travis CI and AppVeyor) are successful
2.  Update documentation whenever making changes to the API or functionality.
3.  Add a note to `CHANGELOG.md` about the changes. Include a link to the
    pull request.

[0]: https://github.com/plus3it/terrastory/issues
[1]: https://travis-ci.org/plus3it/terrastory/pull_requests
