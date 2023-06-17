// pipeline {
//     agent any
//     tools { go '1.20' }
//     environment {
//         GO114MODULE = 'on'
//         CGO_ENABLED = 0 
//         GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
//     }
//     stages {
//         stage("unit-test") {
//             steps {
//                 echo 'UNIT TEST EXECUTION STARTED'
//                 sh 'make unit-tests'
//             }
//         }
//         stage("functional-test") {
//             steps {
//                 echo 'FUNCTIONAL TEST EXECUTION STARTED'
//                 sh 'make functional-tests'
//             }
//         }
//         stage("build") {
//             steps {
//                 echo 'BUILD EXECUTION STARTED'
//                 sh 'go version'
//                 sh 'go get ./...'
//                 sh 'docker build -t learngo:latest .'
//                 sh 'docker run learngo'
//             }
//         }
//         // stage('deliver') {
//         //     agent any
//         //     steps {
//         //         withCredentials([usernamePassword(credentialsId: 'dockerhub', passwordVariable: 'dockerhubPassword', usernameVariable: 'dockerhubUser')]) {
//         //         sh "docker login -u ${env.dockerhubUser} -p ${env.dockerhubPassword}"
//         //         sh 'docker push shadowshotx/product-go-micro'
//         //         }
//         //     }
//         // }
//     }
// }

pipeline {
  // Run on an agent where we want to use Go
//   agent any

  // Ensure the desired Go version is installed for all stages,
  // using the name defined in the Global Tool Configuration
//   tools { go '1.19' }

  stages {
    stage('Build') {
      steps {
        // Output will be something like "go version go1.19 darwin/arm64"
        echo 'go version'
      }
    }
  }
}