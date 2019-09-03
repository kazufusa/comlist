package comlist

import "golang.org/x/sys/windows/registry"

func Coms() (coms []string, err error) {
	key, err := registry.OpenKey(
		registry.LOCAL_MACHINE,
		"HARDWARE\\DEVICEMAP\\SERIALCOMM",
		registry.QUERY_VALUE,
	)
	if err != nil {
		return coms, err
	}
	defer key.Close()

	ks, err := key.Stat()
	if err != nil {
		return coms, err
	}

	coms = make([]string, ks.ValueCount)

	names, err := key.ReadValueNames(int(ks.ValueCount))
	if err != nil {
		return coms, err
	}

	for i, name := range names {
		coms[i], _, err = key.GetStringValue(name)
		if err != nil {
			return coms, err
		}
	}

	return coms, nil
}
