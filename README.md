# fetching-github-contrib
Codes to fetch information about the contribution of a specific repository

# GITHUB Auth 설정 방법

아무런 Auth 설정 없이 GitHub에 API를 호출하면 1시간 당 60건 밖에 요청할 수 없습니다. 하지만 API 요청을 Auth Token과 함께 보내면 1시간 당 5,000건으로 제한이 완화되므로 Auth 관련 설정을 꼭 해야 합니다.

그 방법은 아래와 같습니다.

1. 본 리포지토리의 `/backend` 경로에 `.env` 파일을 만들어주세요.
2. 그리고 그 파일에 `GITHUB_TOKEN = [토큰]` (대괄호 제외)를 적어주세요. 토큰은 자신의 GitHub 계정 설정을 통해 생성해야 합니다.
3. `.env` 파일을 저장하고 서버를 다시 시작해주세요.

이렇게 하면 API 요청을 토큰을 생성한 계정의 권한으로 보내게 되므로 1시간 당 5,000건의 요청을 보낼 수 있습니다.