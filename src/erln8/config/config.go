package config

type Erln8Config struct {
  LinkDir string
}

type OTPDownloadSource struct {
  Name string
  URL string
}

type OTPVersion struct {
  Name string
  Major string
  Minor string
}

type OTPCompilerFlagsSource struct {
  Name string
  URL string
}

type OTPCompileFlags struct {
  Platform string
  Tag string // release vs debug
  Version string
  Flags string
}


