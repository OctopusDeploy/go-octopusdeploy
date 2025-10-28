package machines

import "encoding/json"

type ArgoCDInstallDetails struct {
	GatewayVersion      string `json:"GatewayVersion"`
	HelmChartVersion    string `json:"HelmChartVersion"`
	HelmReleaseName     string `json:"HelmReleaseName"`
	KubernetesNamespace string `json:"KubernetesNamespace"`
	ArgoCDVersion       string `json:"ArgoCDVersion,omitempty"`
}

// MarshalJSON returns a deployment target as its JSON encoding.
func (a *ArgoCDInstallDetails) MarshalJSON() ([]byte, error) {
	installDetails := struct {
		GatewayVersion      string `json:"GatewayVersion"`
		HelmChartVersion    string `json:"HelmChartVersion"`
		HelmReleaseName     string `json:"HelmReleaseName"`
		KubernetesNamespace string `json:"KubernetesNamespace"`
		ArgoCDVersion       string `json:"ArgoCDVersion,omitempty"`
	}{
		GatewayVersion:      a.GatewayVersion,
		HelmChartVersion:    a.HelmChartVersion,
		HelmReleaseName:     a.HelmReleaseName,
		KubernetesNamespace: a.KubernetesNamespace,
		ArgoCDVersion:       a.ArgoCDVersion,
	}

	return json.Marshal(installDetails)
}

// MarshalJSON returns a deployment target as its JSON encoding.
func (resource *ArgoCDInstallDetails) UnmarshalJSON(b []byte) error {
	var argoCDInstallDetails map[string]*json.RawMessage
	err := json.Unmarshal(b, &argoCDInstallDetails)
	if err != nil {
		return err
	}

	for installDetailsKey, installDetailsValue := range argoCDInstallDetails {
		switch installDetailsKey {
		case "GatewayVersion":
			err = json.Unmarshal(*installDetailsValue, &resource.GatewayVersion)
			if err != nil {
				return err
			}
		case "HelmChartVersion":
			err = json.Unmarshal(*installDetailsValue, &resource.HelmChartVersion)
			if err != nil {
				return err
			}
		case "HelmReleaseName":
			err = json.Unmarshal(*installDetailsValue, &resource.HelmReleaseName)
			if err != nil {
				return err
			}
		case "KubernetesNamespace":
			err = json.Unmarshal(*installDetailsValue, &resource.KubernetesNamespace)
			if err != nil {
				return err
			}
		case "ArgoCDVersion":
			err = json.Unmarshal(*installDetailsValue, &resource.ArgoCDVersion)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
