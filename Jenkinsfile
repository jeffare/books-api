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
                    app = docker.build("<DOCKER_HUB_USERNAME>/books-api")
                    app.inside {
                        sh 'echo $(curl localhost:8080)'
                    }
                }
            }
        }
stage('Push Docker Image') {
            when {
                branch 'master'
            }
            steps {
                script {
                    docker.withRegistry('https://registry.hub.docker.com', 'docker_hub_login') {
                        app.push("${env.BUILD_NUMBER}")
                        app.push("latest")
                    }
                }
            }
        }
    }   
}
