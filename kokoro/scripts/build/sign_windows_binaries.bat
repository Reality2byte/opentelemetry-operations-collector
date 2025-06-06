@REM Copyright 2025 Google LLC
@REM
@REM Licensed under the Apache License, Version 2.0 (the "License");
@REM you may not use this file except in compliance with the License.
@REM You may obtain a copy of the License at
@REM
@REM     http://www.apache.org/licenses/LICENSE-2.0
@REM
@REM Unless required by applicable law or agreed to in writing, software
@REM distributed under the License is distributed on an "AS IS" BASIS,
@REM WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
@REM See the License for the specific language governing permissions and
@REM limitations under the License.

@REM Copy input directory into the artifacts directory recursively.
robocopy "%KOKORO_GFILE_DIR%"\dist "%KOKORO_ARTIFACTS_DIR%"\dist /E

ksigntool sign GOOGLE_EXTERNAL /v /debug /t http://timestamp.digicert.com "%KOKORO_ARTIFACTS_DIR%"\dist\otelcol-google-windows_windows_amd64_v1\*.exe

