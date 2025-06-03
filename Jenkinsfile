pipeline {
    agent any
    environment {
        DOCKER_HUB_USERNAME ='jeffare9x'
    }
    stages {
stage('Build Docker Image') {
            when {
                branch 'main'
            }
            steps {
                script {
                    app = docker.build("jeffare9x/books-api")
                }
            }
        }
stage('Push Docker Image') {
            when {
                branch 'main'
            }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker-agent') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }
    }   
}
