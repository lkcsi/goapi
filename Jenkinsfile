pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                sh 'docker stop $(docker ps -a -q)'
                sh 'docker container prune -f'
                sh 'docker-compose up -d'
            }
        }
    }
}
