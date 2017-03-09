Rails.application.config.middleware.use OmniAuth::Builder do
  provider :github, ENV['GITHUB_KEY'], ENV['GITHUB_SECRET'], scope: "user,repo"
  puts "*" * 30
  p ENV['GITHUB_KEY']
  p ENV['GITHUB_SECRET']
  puts "-" * 30
end
