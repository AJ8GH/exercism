package erratum

func Use(opener ResourceOpener, input string) (err error) {
	r, err := openResource(opener)
	if err != nil {
		return err
	}

	defer r.Close()
	defer func() {
		if e := recover(); e != nil {
			f, isFrob := e.(FrobError)
			if isFrob {
				r.Defrob(f.defrobTag)
			}
			err = e.(error)
		}
	}()

	r.Frob(input)
	return err
}

func openResource(opener ResourceOpener) (Resource, error) {
	r, err := opener()

	for err != nil {
		_, isTransient := err.(TransientError)
		if !isTransient {
			return nil, err
		}
		r, err = opener()
	}

	return r, nil
}
