pipeline {
    agent any
    environment {
        AUTH_SECRET=credentials('auth-secret')
        AUTH_ENABLED=auth-enabled
        BOOKS_REPOSITORY='SQL'
        BOOKS_DB_HOST='books-db-1'
        BOOKS_DB_PASSWORD=credentials('books-db-password')
        BOOKS_DB_PORT=3306
        BOOKS_API_PORT=8081
    }
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
                    echo "${BOOKS_API_PORT}"
                    sh 'docker compose build'
                    sh 'docker compose up -d'
                }
            }
        }
    }
}
