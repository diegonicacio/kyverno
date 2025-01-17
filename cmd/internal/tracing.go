package internal

import (
	"context"
	"net"

	"github.com/go-logr/logr"
	"github.com/kyverno/kyverno/pkg/tracing"
	"k8s.io/client-go/kubernetes"
)

func SetupTracing(logger logr.Logger, name string, kubeClient kubernetes.Interface) context.CancelFunc {
	logger = logger.WithName("tracing").WithValues("enabled", tracingEnabled, "address", tracingAddress, "port", tracingPort, "creds", tracingCreds)
	logger.Info("setup tracing...")
	if tracingEnabled {
		shutdown, err := tracing.NewTraceConfig(
			logger,
			name,
			net.JoinHostPort(tracingAddress, tracingPort),
			tracingCreds,
			kubeClient,
		)
		checkError(logger, err, "failed to setup tracing")
		return shutdown
	}
	return func() {}
}
