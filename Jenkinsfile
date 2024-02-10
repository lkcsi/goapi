pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                sh 'docker-compose up'
            }
        }
    }
}