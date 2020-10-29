Pod::Spec.new do |spec|
  spec.name         = 'Gpch'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/psdchaineum/go-psdchaineum'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Psdchain Client'
  spec.source       = { :git => 'https://github.com/psdchaineum/go-psdchaineum.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gpch.framework'

	spec.prepare_command = <<-CMD
    curl https://gpchstore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gpch.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
