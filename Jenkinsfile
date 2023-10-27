pipeline {
    agent any

    stages {
        stage('Build and Test') {
            steps {
                script {
                    sh 'go mod download'
                    sh 'go test -v'
                }
            }
        }

        stage('Generate HTML Coverage Report') {
            steps {
                script {
                    sh 'go test -coverprofile=coverage.out'
                    sh 'go tool cover -html=coverage.out -o coverage.html'
                }
            }
        }

        stage('Archive HTML Coverage Report') {
            steps {
                archiveArtifacts 'coverage.html'
            }
        }

        stage('Build and Upload to Nexus') {
            steps {
                script {
                    // Adjust these commands based on how you build and upload your Go application to Nexus
                    sh 'go build -o your-app'
                    // sh 'curl -u username:password -X PUT --upload-file your-app https://nexus.example.com/repository/your-repo/your-app/1.0.0/your-app-1.0.0'
                }
            }
        }
    }

    post {
        always {
            cleanWs()
        }
    }
}
