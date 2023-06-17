pipeline {
    agent {
        go {
            // Specify the name of the Go installation configured in Global Tool Configuration
            name 'Go'
        }
    }
    stages {
        stage('Build') {
            steps {
                // Add your build steps here
                sh 'go build'
            }
        }
        // Add more stages as needed
    }
}
