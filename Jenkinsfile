pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                catchError {
                    sh 'docker stop $(docker ps -a -q)'
                    sh 'docker rm $(docker ps -a -q)'
                }
            }
        }
    }
}
