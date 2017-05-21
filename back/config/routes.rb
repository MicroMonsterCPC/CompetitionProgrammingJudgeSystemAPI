Rails.application.routes.draw do
  # For details on the DSL available within this file, see http://guides.rubyonrails.org/routing.html
  namespace :api do
    namespace :questions do
      get '/', action: 'index'
    end
  end

  get 'auth/:provider/callback', to: 'sessions#callback'
  get '/logout', to: 'sessions#destroy', as: :logout

  post "/answers/create", to: "answers#create"


  root 'home#top'
  resources :questions
end
