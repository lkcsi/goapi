pipeline {
    agent any
    stages {
        stage("build") {
            steps {
                echo 'build'
                script {
                    try {
                        sh 'docker stop $(docker ps -a -q)'
                        sh 'docker rm $(docker ps -a -q)'
                    } catch (err) {
                        echo 'No probs'
                    }
                    sh 'docker-compose build'
                    sh 'docker-compose up -d'
                }
            }
        }
    }
}
