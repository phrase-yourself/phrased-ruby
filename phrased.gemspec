lib = File.expand_path('../lib', __FILE__)
$LOAD_PATH.unshift(lib) unless $LOAD_PATH.include?(lib)
require 'phrased/version'

Gem::Specification.new do |spec|
  spec.name          = 'phrased'
  spec.version       = Phrased::VERSION
  spec.authors       = ['Jakob Holderbaum']
  spec.email         = ['author@phrased.org']

  spec.summary       = 'Generate awesome passphrases from a set of dictionaries'
  spec.description   = 'Passphrases became widely used through diceware.'
  spec.homepage      = 'https://phrased.org/'
  spec.license       = 'MIT'

  spec.files         = `git ls-files -z`.split("\x0").reject do |f|
    f.match(%r{^(test|spec|features)/})
  end
  spec.require_paths = ['lib']

  spec.add_development_dependency 'bundler', '~> 1.15'
  spec.add_development_dependency 'minitest', '~> 5.10'
  spec.add_development_dependency 'rubocop', '~> 0.51'
end
