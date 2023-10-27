pipeline {
    agent any

   tools {
       go 'go-1.21.3' 
        }

    environment {
        SONARQUBE_SERVER = credentials('SONARQUBE_SERVER') // Reference Jenkins credential ID
    }

    parameters {
        string(name: 'SONAR_PROJECT_KEY', defaultValue: 'default-key', description: 'SonarQube Project Key')
        string(name: 'SONAR_PROJECT_NAME', defaultValue: 'default-name', description: 'SonarQube Project Name')
    }

    stages {

        stage('Generate SonarQube Properties') {
            steps {
                script {
                    def sonarProperties = """
                        sonar.projectKey=${params.SONAR_PROJECT_KEY}
                        sonar.projectName=${params.SONAR_PROJECT_NAME}
                        sonar.sources=./
                        sonar.tests=./
                        sonar.language=go
                    """

                    writeFile file: 'sonar-project.properties', text: sonarProperties
                }
            }
        }

        stage('Unit Test') {
            steps {
                script {
                    sh 'go mod init todo-app'
                    sh 'go test -v'
                }
            }
        }

        stage('Coverage Report') {
            steps {
                script {
                    sh 'go test -coverprofile=coverage.out'
                    sh 'go tool cover -html=coverage.out -o coverage.html'
                }
            }
            steps {
                archiveArtifacts 'coverage.html'
            }
        }

        // stage('Archive HTML Coverage Report') {
        //     steps {
        //         archiveArtifacts 'coverage.html'
        //     }
        // }

        stage('Run SonarQube Analysis') {
            steps {
                script {
                    withSonarQubeEnv('SonarQubeServer') {
                        sh 'sonar-scanner'
                    }
                }
            }
        }

        stage('Build') {
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
