pipeline {
    agent 'docker'

    stages {
        stage('Build Docker') {
            steps {
                docker build -t snippetbox:B'${BUILD_NUMBER}' -f Dockerfile .
            }
        }
    }
}
