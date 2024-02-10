pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                sh 'docker-compose down'
                sh 'docker-compose up -d'
            }
        }
    }
}
