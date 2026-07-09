# =============================================================================
# Golden Base Image — 조직 표준 런타임 베이스
# =============================================================================
# 원칙: digest pin (재현성) / distroless (최소 공격표면) / nonroot (권한최소)
# 이 이미지는 Trivy 스캔(HIGH/CRITICAL 차단) 통과 후에만 ECR에 push됨.
# 모든 앱은 이 이미지만 FROM으로 참조.
# =============================================================================

FROM gcr.io/distroless/static@sha256:d5f030ca7c5793784e9ea4178a116da360250411d13921a5af27c6cb5a5949bf

LABEL org.opencontainers.image.title="golden-base"
LABEL org.opencontainers.image.description="Organization-approved base image. Scanned and signed."
LABEL org.opencontainers.image.vendor="Platform Engineering"

USER 65532:65532
