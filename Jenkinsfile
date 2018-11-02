pipeline {
    agent {
        dockerfile {
            filename 'Dockerfile'
        }

    }
    stages {
        stage('Set env') {
            steps {
                sh '''go env
                export GOPATH=$WORKSPACE
                export GOBIN=$GOPATH/bin
                export PATH=$GOPATH:$GOBIN:$PATH
                echo "echo path"
                echo $GOPATH
                #cat /etc/os-release'''
            }
        }
        stage('Tnstall package') {
            steps {
                sh '''go env
                    export GOPATH=$WORKSPACE
                    export GOBIN=$GOPATH/bin
                    export PATH=$GOPATH:$GOBIN:$PATH

                    echo "echo path"
                    echo $GOPATH
                    echo "install package"
                    go get -u github.com/golang/dep/cmd/dep
                    #cd src/agent
                    #dep ensure'''
            }
        }
        stage('Test') {
            steps {
                sh '''export GOPATH=$WORKSPACE
                    export GOBIN=$GOPATH/bin
                    export PATH=$GOPATH:$GOBIN:$PATH
                    #cd src/agent
                    go test -coverprofile="coverage.out" ./...
                    go test -json > report.json ./...
                '''
            }
        }
        #stage('sonarqube') {
        #    steps {
        #        echo 'sanner'
        #        withSonarQubeEnv('Sonar') {
        #            sh 'mvn org.sonarsource.scanner.maven:sonar-maven-plugin:3.3.0.603:sonar ' +
        #             '-f all/pom.xml ' +
        #             '-Dsonar.host.url=http://172.18.0.3:9000' +
        #             '-Dsonar.login=agent ' +
        #             '-Dsonar.password=agent ' +
        #             '-Dsonar.projectKey=agent-connector' +
        #             '-Dsonar.projectName=agent-connector' +
        #             '-Dsonar.projectVersion=1.1.1' +
        #             '-Dsonar.language=go ' +
        #             '-Dsonar.sourceEncoding=UTF-8 ' +
        #             '-Dsonar.sources=src/agent ' +
        #             '-Dsonar.go.tests.reportPaths=src/agent/report.json ' +
        #             '-Dsonar.go.coverage.reportPaths=src/agent/coverage.out '
        #        }
        #    }
        #}
        stage('artifacts') {
            steps {
                echo 'save'
            }
        }
    }
}