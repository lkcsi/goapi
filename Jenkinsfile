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
                    } catch {
                        echo 'No probs'
                    }
                }
            }
        }
    }
}
