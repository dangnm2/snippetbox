node('docker') {
    stage 'Checkout'
        checkout scm
    stage 'Build & UnitTest'
    sh "docker build -t snippetbox:B${BUILD_NUMBER} -f Dockerfile ."
}