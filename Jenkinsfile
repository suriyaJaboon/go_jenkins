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
                    #cd $WORKSPACE/Go_SLG
	            echo $PWD
		    go run main.go
                    #dep ensure'''
            }
        }
        stage('Test') {
            steps {
                sh '''
		    export GOPATH=$WORKSPACE
                    export GOBIN=$GOPATH/bin
                    export PATH=$GOPATH:$GOBIN:$PATH
		    echo $PWD
		    go test sum_test.go sum.go
                    #cd Go_SLG
                    #go test -coverprofile="coverage.out" ./...
                    #go test -json > report.json ./...
                '''
            }
        }
	stage('Build Production'){	
            steps {
                sh '''
		    export GOPATH=$WORKSPACE
                    export GOBIN=$GOPATH/bin
                    export PATH=$GOPATH:$GOBIN:$PATH
		    echo $PWD
		    go build -o /var/jenkins_home/go_slg main.go
		    cd /var/jenkins_home && ./go_slg
                    #cd Go_SLG
                    #go test -coverprofile="coverage.out" ./...
                    #go test -json > report.json ./...
                '''
            }
        }
	
        stage('artifacts') {
            steps {
                echo 'save'
            }
        }
    }
}
