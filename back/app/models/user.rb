class User < ApplicationRecord
  has_many :answers
  has_many :questions

  def self.create_with_omniauth(auth)
    user = User.new
    user.provider = auth['provider']
    user.uid = auth['uid']
    user.name = auth['info']['nickname']
    user.oauth_token = auth['credentials']['token']
    user.save
    return user
  end
end
