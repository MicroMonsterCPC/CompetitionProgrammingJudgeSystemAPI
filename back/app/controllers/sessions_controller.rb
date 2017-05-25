class SessionsController < ApplicationController
  def callback
    auth = request.env['omniauth.auth']
    user = User.find_by(provider: auth['provider'], uid: auth['uid'])
    unless user
      user = User.create_with_omniauth(auth)
    end
    puts "*" * 50
    p user
    puts "*" * 50
    session[:user_id] = user.id
    redirect_to root_path
  end

  def destroy
    reset_session
    redirect_to root_path
  end
end
