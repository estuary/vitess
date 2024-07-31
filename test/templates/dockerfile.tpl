ARG bootstrap_version=33
ARG image="vitess/bootstrap:${bootstrap_version}-{{.Platform}}"

FROM "${image}"

USER root

# Re-copy sources from working tree
RUN rm -rf /vt/src/github.com/estuary/vitess/*
COPY . /vt/src/github.com/estuary/vitess

{{if .InstallXtraBackup}}
# install XtraBackup
RUN wget https://repo.percona.com/apt/percona-release_latest.$(lsb_release -sc)_all.deb
RUN apt-get update
RUN apt-get install -y gnupg2
RUN dpkg -i percona-release_latest.$(lsb_release -sc)_all.deb
RUN apt-get update
RUN apt-get install -y percona-xtrabackup-24
{{end}}

# Set the working directory
WORKDIR /vt/src/github.com/estuary/vitess

# Fix permissions
RUN chown -R vitess:vitess /vt

USER vitess

# Set environment variables
ENV VTROOT /vt/src/github.com/estuary/vitess
# Set the vtdataroot such that it uses the volume mount
ENV VTDATAROOT /vt/vtdataroot

# create the vtdataroot directory
RUN mkdir -p $VTDATAROOT

# install goimports
RUN go install golang.org/x/tools/cmd/goimports@latest

{{if .MakeTools}}
# make tools
RUN make tools
{{end}}

# sleep for 50 minutes
CMD sleep 3000
