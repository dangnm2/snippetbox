pipeline {
    agent 'docker'

    stages {
        stage('Build Docker') {
            steps {
                sh "docker build -t snippetbox:B'${BUILD_NUMBER}' -f Dockerfile ."
            }
        }
    }
}
