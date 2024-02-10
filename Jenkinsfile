pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                sh 'docker stop $(docker ps -a -q)'
                sh 'docker prune -y'
                sh 'docker-compose up'
            }
        }
    }
}
