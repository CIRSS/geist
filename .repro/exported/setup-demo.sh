
sudo apt update
sudo apt -y install graphviz
sudo apt -y install python3 python3-pip
pip3 install pandas

# Download Apache Jena and add arq to PATH
jena_version=3.17.0
jena_version_name=apache-jena-${jena_version}
jena_archive=${jena_version_name}.tar.gz

cd ${HOME}
wget http://archive.apache.org/dist/jena/binaries/${jena_archive}
tar xvvf ${jena_archive}
rm ${jena_archive}
 
repro.setenv JENA_HOME ${HOME}/${jena_version_name}
repro.setenv JENA_CLASSPATH '${JENA_HOME}/lib/*'

repro.prefixpath '${JENA_HOME}/bin'
