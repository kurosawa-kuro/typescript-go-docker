FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install Node.js and Japanese language support
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
        nodejs \
        locales \
    && sed -i '/ja_JP.UTF-8/s/^# //g' /etc/locale.gen \
    && locale-gen \
    && echo "LANG=ja_JP.UTF-8" > /etc/default/locale \
    && echo "LANGUAGE=ja_JP:ja" >> /etc/default/locale \
    && echo "LC_ALL=ja_JP.UTF-8" >> /etc/default/locale

# Set environment variables for Japanese locale
ENV LANG=ja_JP.UTF-8 \
    LANGUAGE=ja_JP:ja \
    LC_ALL=ja_JP.UTF-8

# Install global npm packages if needed
# RUN su vscode -c "npm install -g typescript @types/node" 2>&1

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>

# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
