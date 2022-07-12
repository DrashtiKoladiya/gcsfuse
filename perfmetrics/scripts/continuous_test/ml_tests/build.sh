set -e
cd "${KOKORO_ARTIFACTS_DIR}/github/gcsfuse/perfmetrics/scripts"
echo "Running tests"
cd ml_tests
source setup.sh