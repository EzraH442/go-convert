pipeline {
  agent any

  stages {
    stage('build main') {
      steps {
        sh 'docker build -t ezraweb/go-convert:latest .'
      }
    }
  }
}